import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "./components";
import Login from "./pages/login";
import Projects from "./pages/dashboard/projects/index";
import ProjectDetail from "./pages/dashboard/projects/ProjectDetail";
import Users from "./pages/dashboard/users";
import RequireAuth from "./utils/RequireAuth";

function App() {
	return (
		<Routes>
			<Route path='/login' element={<Login />} />
			<Route element={<Layout />}>
				<Route element={<RequireAuth />}>
					<Route path='/' element={<h1>Hello Home!</h1>} />
					<Route path='/projects'>
						<Route index element={<Projects />} />
						<Route path=':id' element={<ProjectDetail />} />
					</Route>
					<Route path='/users' element={<Users />} />
				</Route>
			</Route>
		</Routes>
	);
}

export default App;
