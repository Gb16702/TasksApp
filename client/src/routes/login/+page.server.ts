import { fail, type Actions, redirect } from "@sveltejs/kit";

export const actions: Actions = {
  login: async ({ request, cookies }) => {
    const formData = await request.formData();

    const user = Object.fromEntries(formData);

    const errors: Record<string, unknown> = {};

    if (!user.email || typeof user.email !== "string") {
      errors.email = "Adresse mail invalide";
    }

    if (!user.password || typeof user.password !== "string") {
      errors.password = "Mot de passe invalide";
    }

    if (Object.keys(errors).length > 0) {
      const data = {
        data: Object.fromEntries(formData),
        errors,
      };

      return fail(400, data);
    }

    const response = await fetch("http://127.0.0.1:8000/api/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: user.email,
        password: user.password,
      }),
    });

    const data = await response.json();

    cookies.set("secure__session", data.session, {
      secure: false,
      httpOnly: true,
      maxAge: 60 * 60 * 24 * 7,
      path: "/",
      });

    cookies.set("uid", data.ID, {
      secure: false,
      httpOnly: true,
      maxAge: 60 * 60 * 24 * 7,
      path: "/",
    })

    throw redirect(300, "/");
  },
};
