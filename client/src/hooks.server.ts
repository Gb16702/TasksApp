import { redirect, type Handle } from '@sveltejs/kit';
import decodeToken from './utils/decodeToken';

export const handle: Handle = async ({ event, resolve }) => {

	let session;
	if (event.request.headers.get('cookie')) {

	 	const sessionCookieValue: string | undefined = event.request.headers
	 		.get('cookie')?.split(";").find((row) => row.startsWith("secure__session"))?.substring(16)
	 	if (sessionCookieValue) {
	 		session = await decodeToken(sessionCookieValue);
	 	}
	}

	if ((event.url.pathname == '/register') || (event.url.pathname == '/login')) {
		if (session) {
			throw redirect(303, '/');
		}
	}

	if (event.url.pathname == "/") {
		if(!session) {
			throw redirect(303, '/login');
		}
	}

	const response = await resolve(event);
	return response;
};
