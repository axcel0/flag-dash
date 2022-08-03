import axios from "axios";
import RefreshToken from "./RefreshToken";

const APIClient = () => {
	const instance = axios.create();
	instance.interceptors.request.use(async (req) => {
		const token = localStorage.getItem("token");
		if (token) {
			req.headers = {
				...req.headers,
				Authorization: token,
			};
		}
		return req;
	});

	instance.interceptors.response.use(
		async (res) => {
			return res;
		},
		async (err) => {
			const config = err?.config;

			if (err.response.status === 401 && !config.sent) {
				const result = await RefreshToken();

				if (result.token) {
					config.headers = {
						...config.headers,
						Authorization: result.token,
					};
				}
				localStorage.setItem("token", result.token);
				return instance(config);
			}
			return Promise.reject(err);
		},
	);

	return instance;
};

export default APIClient();
