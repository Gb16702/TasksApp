import { redirect } from '@sveltejs/kit';
import decodeToken from '../../../utils/decodeToken';
import type { PageServerLoad } from '../$types';

export const load: PageServerLoad = async ({ url, cookies }: any) => {
	const trimmedUrl: string = url.pathname.split('/')[2].slice(0, url.pathname.length - 1);

	const value = await decodeToken(trimmedUrl);

	const errors: Record<string, unknown> = {};

	const isExpired: boolean = Date.now() >= value.exp * 1000;

	if (isExpired) {
		 errors.token = 'Le lien a expiré';
		 throw redirect(303, "/register")
	}

	const response = await fetch('http://127.0.0.1:8000/api/auth/verify-email', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			email: value.user,
			token: trimmedUrl
		})
	});

	if (!response.ok) {
		errors.token = 'Le lien est invalide';
		throw redirect(303, "/register")
	}

	const data = await response.json();

	const sessionResponse: Response = await fetch('http://127.0.0.1:8000/api/auth/login', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			"Authorization": `Bearer ${trimmedUrl}`
		}});

	const sessionData = await sessionResponse.json();

	console.log(sessionData);

	const session = sessionData.data;

	cookies.set("secure__session", session, {
		secure: false,
		httpOnly: true,
		maxAge: 60 * 60 * 24 * 7,
		path: '/'
	});

	return {
		props: {
			data,
			errors
		}
	};
};
