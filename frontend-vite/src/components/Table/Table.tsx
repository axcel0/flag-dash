import React from "react";

interface TableProps {
	heads: string[];
	values: any[];
}

const Table: React.FC<TableProps> = ({ heads, values }) => {
	return (
		<table className='table-auto text-left'>
			<thead>
				<tr>
					{heads.map((head, index) => (
						<th key={index} className='px-4 py-2'>
							{head}
						</th>
					))}
				</tr>
			</thead>
			<tbody>
				{values.map((value, index) => (
					<tr key={index}>
						{Object.entries(value).map(([k, v]) => (
							<td key={k} className='px-4 py-2'>
								{v as any}
							</td>
						))}
					</tr>
				))}
			</tbody>
		</table>
	);
};

export default Table;
