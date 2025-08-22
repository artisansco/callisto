import { redirect } from "@sveltejs/kit";
import { z } from "zod";
import { form, getRequestEvent } from "$app/server";

const register_schema = z
  .object({
    name: z.string().trim().min(2, { error: "Name must be at least 2 characters long" }),
    email: z.email({ error: "Please enter a valid email address" }),
    password: z.string().min(4, { error: "Password must be at least 4 characters long" }),
    confirmPassword: z
      .string()
      .min(4, { error: "Confirm Password must be at least 4 characters long" }),
  })
  .refine(({ password, confirmPassword }) => password === confirmPassword, {
    error: "Passwords do not match",
  });

export const register = form(async (form_data) => {
  const form = Object.fromEntries(form_data);
  const { success, error } = register_schema.safeParse(form);

  if (!success) {
    const message = error.issues.at(0)?.message;
    return { error: message };
  }

  try {
    await new Promise((resolve) => setTimeout(resolve, 1000));

    const { cookies } = getRequestEvent();
    cookies.set("token", "mock-new-user-session-token", {
      path: "/",
      httpOnly: true,
      secure: true,
      sameSite: "strict",
      maxAge: 60 * 60 * 24, // 1 day
    });

    // Redirect to dashboard
    redirect(307, "/dashboard");
  } catch (_e) {
    // @ts-expect-error
    return { error: _e.message };
  }
});
