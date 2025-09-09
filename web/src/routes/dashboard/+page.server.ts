import { redirect } from "@sveltejs/kit";
// import { getUser, requireAuth } from "$lib/server/auth";

/** @type {import('./$types').PageServerLoad} */
export async function load({ cookies }) {
	// Get user data for the authenticated user
	// const user = getUser(cookies);

	// Return user data to the page
	return {
		user: { id: "01992ee0-16da-7831-812e-9471bcce8fdb", name: "John Doe" },
	};
}
