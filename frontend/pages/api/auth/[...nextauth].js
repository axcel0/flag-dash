import axios from "axios";
import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export default NextAuth({
	providers: [
		CredentialsProvider({
			name: "Golang Auth API",
			credentials: {
				email: {
					label: "Email",
					type: "text",
					placeholder: "text@example.com",
				},
				password: { label: "Password", type: "password" },
			},
			async authorize(credentials, req) {
				const res = await axios.post(
					"http://127.0.0.1:3001/api/v1/auth/login",
					{
						email: credentials.email,
						password: credentials.password,
					},
				);
				if (res) {
					return res.data;
				} else {
					return null;
				}
			},
		}),
	],
	callbacks: {
		async jwt({ token, user, account, profile, isNewUser }) {
			return user;
		},
	},
});
