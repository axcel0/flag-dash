import { useFormik } from "formik";

import { FormProps } from ".";

import {
	useGetFlagQuery,
	useEditFlagMutation,
} from "../../redux/features/flags/flagApiSlice";
import CircularLoading from "../CircularLoading/CircularLoading";
import ToggleButton from "../ToggleButton/ToggleButton";

const EditFlagForm: React.FC<FormProps> = ({
	handleSuccess,
	handleFailed,
	handleCancel,
	data,
}) => {
	const {
		data: flagData,
		isLoading,
		isSuccess,
	} = useGetFlagQuery<any>({
		id: data?.id,
	});

	const [
		editFlag,
		{ isSuccess: isSuccessEditFlag, isLoading: isLoadingEdit },
	] = useEditFlagMutation();

	const editFlagForm = useFormik({
		enableReinitialize: true,
		initialValues: {
			id: flagData?.flag?.id || -1,
			name: flagData?.flag?.name || "",
			active: flagData?.flag?.active || false,
		},
		onSubmit: async (values) => {
			try {
				const res = await editFlag({
					id: values.id,
					name: values.name,
					active: values.active,
				});
				handleSuccess();
			} catch (err) {
				console.log(err);
				handleFailed();
			}
		},
	});

	if (isLoading || isLoadingEdit) return <CircularLoading />;

	return (
		<form onSubmit={editFlagForm.handleSubmit}>
			<h3 className='text-3xl font-medium'>Edit Flag</h3>
			<h3 className='text-lg'>Name:</h3>
			<input
				className='rounded-md border p-2 w-[400px]'
				type='text'
				name='name'
				onChange={editFlagForm.handleChange}
				value={editFlagForm.values.name}
			/>
			<h3 className='text-lg'>Active:</h3>
			<ToggleButton
				name='active'
				checked={editFlagForm.values.active}
				onChange={editFlagForm.handleChange}
			/>
			<div className='my-2'>
				<button
					className='bg-yellow-300 hover:bg-yellow-200 rounded-md shadow-md p-2 mr-2'
					type='submit'
				>
					<p className='text-lg'>Edit Flag</p>
				</button>
				<button
					className='bg-red-300 hover:bg-red-200 rounded-md shadow-md p-2 ml-2'
					onClick={() => handleCancel()}
				>
					<p className='text-lg'>Cancel</p>
				</button>
			</div>
		</form>
	);
};

export default EditFlagForm;
