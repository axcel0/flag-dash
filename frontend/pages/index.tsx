import type { ReactElement } from "react";
import type { NextPage } from "next";
import { signIn } from "next-auth/react";
import Layout from "../components/Layout";
import type { NextPageWithLayout } from "./_app";

const Home: NextPageWithLayout = () => {
	return (
		<>
			<div className='p-5'>
				<h1 className='text-5xl font-bold'>
					Welcome to Flag Dashboard!
				</h1>
			</div>
		</>
	);
};

Home.getLayout = function getLayout(page: ReactElement) {
	return <Layout>{page}</Layout>;
};

export default Home;
