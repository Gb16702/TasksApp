import { redirect } from "@sveltejs/kit";
import decodeToken from "../../../utils/decodeToken";
import type { PageServerLoad } from "../$types";
import { SERVER_URL } from "$env/static/private";

export const load: PageServerLoad = async ({ url, cookies }: any) => {
  const trimmedUrl: string = url.pathname
    .split("/")[2]
    .slice(0, url.pathname.length - 1);

  const value = await decodeToken(trimmedUrl);

  const errors: Record<string, unknown> = {};

  const isExpired: boolean = Date.now() >= value.exp * 1000;

  if (isExpired) {
    errors.token = "Le lien a expiré";
    throw redirect(303, "/register");
  }

  const response = await fetch(SERVER_URL +"/api/auth/verify-email", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: value.user,
      token: trimmedUrl,
    }),
  });

  if (!response.ok) {
    errors.token = "Le lien est invalide";
    throw redirect(303, "/register");
  }

  const data = await response.json();

  const sessionResponse: Response = await fetch(
    SERVER_URL +"/api/auth/login",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${trimmedUrl}`,
      },
    }
  );

  const sessionData = await sessionResponse.json();

  const session = sessionData.session;

  cookies.set("secure__session", session, {
    secure: false,
    httpOnly: true,
    maxAge: 60 * 60 * 24 * 7,
    path: "/",
  });

  cookies.set("uid", sessionData.ID, {
    secure: false,
    httpOnly: true,
    maxAge: 60 * 60 * 24 * 7,
    path: "/",
  });

  return {
    props: {
      data,
      errors,
    },
  };
};
