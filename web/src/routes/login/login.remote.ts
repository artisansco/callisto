import { redirect } from "@sveltejs/kit";
import z from "zod";
import { form, getRequestEvent } from "$app/server";
import { API_ENDPOINT } from "$env/static/private";

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
	const { success, data:parsed, error } = login_schema.safeParse(form);

	if (!success) {
		const message = error.issues.at(0)?.message;
		return { error: message };
	}

	const { cookies } = getRequestEvent();
	try {
	const res = await fetch(`${API_ENDPOINT}/api/v1/auth/token`, {
		method: "POST",
		headers: {			"Content-Type": "application/json"},
		body: JSON.stringify(parsed),
	});
	const {message,data} = await res.json();

	if (!res.ok) {
		return { error: message };
	}

	cookies.set("token", data.token, {
		path: "/",
		httpOnly: true,
		secure: true,
		sameSite: "strict",
		// maxAge: data.remember ? 60 * 60 * 24 * 7 : 60 * 60 * 24, // 30 days if remember, 1 day otherwise
		maxAge: 60 * 60 * 24, // 30 days if remember, 1 day otherwise
	});

	} catch (_e) {
		// @ts-expect-error
		return { error: _e.message };
	}

	redirect(308, "/dashboard");
});
