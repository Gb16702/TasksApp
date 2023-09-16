import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import testPattern from '../../utils/patternTest';
import {SERVER_URL} from "$env/static/private"


const sleep = async (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

export const actions: Actions = {
	register: async ({ request }) => {
		const formData = await request.formData();

		const user = Object.fromEntries(formData);

		const errors: Record<string, unknown> = {};

		((email: string | false) => {
			(!user.email || (email && !testPattern(email.toString()))) &&
				(errors.email = "Format d'adresse mail invalide");
		})(typeof user.email == 'string' ? user.email.trim() : false);

		if (!user.password || typeof user.password !== 'string') {
			errors.password = 'Mot de passe invalide';
		}

		if (!user.passwordConfirm || typeof user.passwordConfirm !== 'string') {
			errors.passwordConfirm = 'Confirmation de mot de passe invalide';
		}

		if (user.password !== user.passwordConfirm) {
			errors.passwordConfirm = 'Les mots de passe ne correspondent pas';
		}

		if (Object.keys(errors).length > 0) {
			const data = {
				data: Object.fromEntries(formData),
				errors
			};
			return fail(400, data);
		}

		const response = await fetch(SERVER_URL + '/api/auth/register', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				email: user.email,
				password: user.password,
				passwordConfirm: user.passwordConfirm
			})
		});

		const data = await response.json();

		if (response.status != 200) {
			return fail(400, data.message);
		  } else {
			return {
			  status: 200,
			  body: data.message
			};
		}
	}
};
