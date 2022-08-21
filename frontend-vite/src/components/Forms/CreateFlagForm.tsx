import { useFormik } from "formik";

import { useNewFlagMutation } from "../../redux/features/flags/flagApiSlice";

// Components
import { ToggleButton } from "../index";

// Interfaces
import { FormProps } from ".";

const CreateFlagForm: React.FC<FormProps> = ({
	handleSuccess,
	handleFailed,
	handleCancel,
	data,
}) => {
	const [newFlag, { isLoading: loadingNewFlag }] = useNewFlagMutation();
	const addCtxForm = useFormik({
		enableReinitialize: true,
		initialValues: {
			projectId: data,
			name: "",
			active: false,
			value: "",
		},
		onSubmit: async (values) => {
			try {
				console.log("Submission:", values);
				const res = await newFlag({
					projectId: values.projectId,
					name: values.name,
					active: values.active,
				}).unwrap();
				if (res.flag) {
					handleSuccess();
				}
			} catch (err: any) {
				if (err.status == 400) {
					handleFailed();
				}
			}
		},
	});
	return (
		<form onSubmit={addCtxForm.handleSubmit}>
			<h3 className='text-4xl my-2'>Create Flag</h3>
			<input
				name='name'
				type='text'
				placeholder='Flag name'
				onChange={addCtxForm.handleChange}
				value={addCtxForm.values.name}
				className='my-2 p-2 border-4 rounded-md w-full'
			/>
			<h3 className='text-2xl'>Active: </h3>
			<div className='flex flex-row items-center'>
				<p className='text-2xl'>Contexts</p>
				<button className='shadow-xl rounded p-2 mx-2 bg-orange-300 hover:bg-orange-200'>
					<p className='text-2xl'>+</p>
				</button>
			</div>
			<div className='max-h-[300px] overflow-y-auto'>
				<div className='rounded shadow border my-2'>
					<input
						type='text'
						placeholder='Name'
						className='p-2 rounded-md border mx-2'
					/>
					<input
						type='text'
						placeholder='Condition'
						className='p-2 rounded-md border mx-2'
					/>
					<input
						type='text'
						placeholder='Value'
						className='p-2 rounded-md border mx-2'
					/>
				</div>
			</div>
			<div className='flex flex-col'>
				<button
					className='shadow rounded-md p-2 bg-green-300 hover:bg-green-200 my-2'
					type='submit'
				>
					<p className='text-xl'>Create</p>
				</button>
				<button
					onClick={() => handleCancel()}
					className='shadow rounded-md p-2 bg-red-300 hover:bg-red-200 my-2'
				>
					<p className='text-xl'>Cancel</p>
				</button>
			</div>
		</form>
	);
};

export default CreateFlagForm;
