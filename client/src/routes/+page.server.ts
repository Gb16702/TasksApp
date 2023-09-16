import { redirect, type Actions, fail } from "@sveltejs/kit";
import decodeToken from "../utils/decodeToken";
import type { PageServerLoad } from "./$types";
import {SERVER_URL} from "$env/static/private"

export const load: PageServerLoad = async ({ cookies }) => {
  const id = cookies.get("uid");

  const response = await fetch(`${SERVER_URL}/api/tasks/${id}`);

  const { tasks } = await response.json();

  tasks.sort((a: any, b: any) => a.id - b.id);

  return {
    tasks,
    id,
  };
};

export const actions = {
  addTodo: async ({ request, cookies }) => {
    const formData = await request.formData();

    const sessionCookie: string | undefined = cookies.get("secure__session");

    let id;
    if (sessionCookie) {
      const { id: sessionId } = await decodeToken(sessionCookie);
      id = sessionId;
    }

    const { name } = Object.fromEntries(formData);

    const errors: Record<string, unknown> = {};

    if (!name || typeof name !== "string") {
      errors.name = "Titre invalide";
    }

    if (Object.keys(errors).length) {
      return {
        status: 400,
        body: errors,
      };
    }

    const response = await fetch(SERVER_URL + "/api/tasks", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name,
        uid: id,
      }),
    });

    const data = await response.json();

    if (response.status != 200) {
      return fail(400, data.message);
    } else {
      return {
        status: 200,
        body: data.message,
      };
    }
  },

  deleteTodo: async ({ request }) => {
    const formData = await request.formData();

    const { todo } = Object.fromEntries(formData);

    const response = await fetch(`${SERVER_URL}/api/tasks/${todo}`, {
      method: "DELETE",
    });

    const data = await response.json();

    if (response.status != 200) {
      return fail(400, data.message);
    } else {
      return {
        status: 200,
        body: data.message,
      };
    }
  },

  toggleStatus: async ({ request }) => {
    const formData = await request.formData();
    const { todo } = Object.fromEntries(formData);

    const response = await fetch(`${SERVER_URL}/api/tasks/${todo}`, {
      method: "PATCH",
    });

    const data = await response.json();

    if (response.status != 200) {
      return fail(400, data.message);
    } else {
      return {
        status: 200,
        body: data.message,
      };
    }
  },
  editName: async ({ request }) => {
    const formData = await request.formData();

    const object: Record<string, unknown> = Object.fromEntries(formData);

    let id: number = Number(object.id);

    const response = await fetch(`${SERVER_URL}/api/tasks/name`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id,
        name: object.name,
      }),
    });

    const data = await response.json();

    if (response.status != 200) {
      return fail(400, data.message);
    } else {
      return {
        status: 200,
        body: data.message,
      };
    }
  },

  clear: async ({ request }) => {
    const formData = await request.formData();

    const object = Object.fromEntries(formData);

    const response = await fetch(
      `${SERVER_URL}/api/tasks/clear/${Number(object.userId)}`,
      {
        method: "DELETE",
      }
    );

    const data = await response.json();

    if (response.status != 200) {
      return fail(400, data.message);
    } else {
      return {
        status: 200,
        body: data.message,
      };
    }
  },

  logout: async (event) => {
    event.cookies.delete("secure__session");
    event.cookies.delete("uid");

    throw redirect(303, "/login");
  },
} satisfies Actions;
