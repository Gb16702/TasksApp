import type { Actions } from "@sveltejs/kit";
import decodeToken from "../utils/decodeToken";
import type { PageServerLoad, RequestEvent } from "./$types";

export const load: PageServerLoad = async ({ cookies }) => {
  const id = cookies.get("uid");

  const response = await fetch(`http://127.0.0.1:8000/api/tasks/${id}`);

  const { tasks } = await response.json();

  tasks.sort((a: any, b: any) => a.id - b.id);

  return {
    tasks,
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

    const response = await fetch("http://127.0.0.1:8000/api/tasks", {
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

    console.log(data);
  },

  deleteTodo: async ({ request }) => {
    const formData = await request.formData();

    const { todo } = Object.fromEntries(formData);

    const response = await fetch(`http://localhost:8000/api/tasks/${todo}`, {
      method: "DELETE",
    });

    const data = await response.json();

    console.log(data);
  },

  toggleStatus: async ({ request }) => {
    const formData = await request.formData();
    const { todo } = Object.fromEntries(formData);

    const response = await fetch(`http://localhost:8000/api/tasks/${todo}`, {
      method: "PATCH",
    });

    const data = await response.json();

    return data.task.Done;
  },
  editName: async ({ request }) => {
    const formData = await request.formData();

    const data: Record<string, unknown> = Object.fromEntries(formData);

    let id: number = Number(data.id);

    const response = await fetch("http://127.0.0.1:8000/api/tasks/name", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id,
        name: data.name,
      }),
    });
  },
} satisfies Actions;
