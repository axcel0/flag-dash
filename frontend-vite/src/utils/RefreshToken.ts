import APIClient from "./APIClient";

const RefreshToken = async () => {
	const refreshToken = localStorage.getItem("refreshToken");
	try {
		const res = await APIClient.post(
			"http://127.0.0.1:3001/api/v1/auth/refresh-token",
			{ refresh_token: refreshToken },
		);

		if (!res.data.token) {
			localStorage.removeItem("token");
			localStorage.removeItem("refreshToken");
		}

		return res.data;
	} catch (err) {
		localStorage.removeItem("token");
		localStorage.removeItem("refreshToken");
	}
};

export default RefreshToken;
