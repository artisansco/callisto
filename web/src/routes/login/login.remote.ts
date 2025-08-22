import { redirect } from "@sveltejs/kit";
import z from "zod";
import { form, getRequestEvent } from "$app/server";

const login_schema = z.object({
	email: z.email({ error: "Please enter a valid email address" }),
	password: z.string().min(4, { error: "Password must be at least 4 characters long" }),
	remember: z
		.string()
		.transform((val) => val === "on")
		.default(false)
		.optional(),
});

export const login = form(async (form_data) => {
	const form = Object.fromEntries(form_data.entries());
	const { success, data, error } = login_schema.safeParse(form);

	if (!success) {
		const message = error.issues.at(0)?.message;
		return { error: message };
	}

	try {
		await new Promise((resolve) => setTimeout(resolve, 750));

		const { cookies } = getRequestEvent();
		cookies.set("token", "mock-session-token", {
			path: "/",
			httpOnly: true,
			secure: true,
			sameSite: "strict",
			maxAge: data.remember ? 60 * 60 * 24 * 7 : 60 * 60 * 24, // 30 days if remember, 1 day otherwise
		});

		// Redirect to dashboard
		redirect(303, "/dashboard");
	} catch (_e) {
		// @ts-expect-error
		return { error: _e.message };
	}
});
