import React from "react";

interface TableProps {
	columnDef: any[];
	values?: any[];
}

const Table: React.FC<TableProps> = ({ columnDef, values }) => {
	return (
		<table className='table-auto w-full'>
			<thead>
				<tr className='text-left'>
					{columnDef.map((col, index) => (
						<th
							key={index}
							className={`w-[${col.colSize}px]`}
						>
							{col.header}
						</th>
					))}
				</tr>
			</thead>
			<tbody>
				{values?.map((value, index) => (
					<tr key={index}>
						{columnDef.map((col, index) => (
							<td key={index}>{col.cell(value)}</td>
						))}
					</tr>
				))}
			</tbody>
		</table>
	);
};

export default Table;
