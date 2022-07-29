import React, { PropsWithChildren, useState } from "react";

import {
	createColumnHelper,
	flexRender,
	getCoreRowModel,
	useReactTable,
	ColumnResizeMode,
	TableOptions,
	ColumnDef,
} from "@tanstack/react-table";

interface TableProps {
	tableData: any;
	columns: any[];
}

const NewTable: React.FC<PropsWithChildren<TableProps>> = ({
	tableData,
	columns,
}) => {
	const [data, setData] = useState(() => [...tableData]);
	const [columnResizeMode, setColumnResizeMode] =
		useState<ColumnResizeMode>("onChange");

	const table = useReactTable({
		data,
		columns,
		columnResizeMode,
		getCoreRowModel: getCoreRowModel(),
	});

	return (
		<div className='overflow-x-auto'>
			<table
				{...{
					className: "rounded-md border-2",
					style: {
						width: table.getTotalSize(),
					},
				}}
			>
				<thead>
					{table.getHeaderGroups().map((headerGroup, index) => (
						<tr key={headerGroup.id}>
							{headerGroup.headers.map((header, index) => (
								<th
									{...{
										key: header.id,
										className:
											"text-left relative border-r-2",
										colSpan: header.colSpan,
										style: {
											width: header.getSize(),
										},
									}}
								>
									{header.isPlaceholder
										? null
										: flexRender(
												header.column
													.columnDef
													.header,
												header.getContext(),
										  )}
									<div
										{...{
											onMouseDown:
												header.getResizeHandler(),
											onTouchStart:
												header.getResizeHandler(),
											className: `absolute top-0 right-0 w-[10px] hover:bg-gray-200 h-full`,
										}}
									/>
								</th>
							))}
						</tr>
					))}
				</thead>
				<tbody>
					{table.getRowModel().rows.map((row, index) => (
						<tr key={row.id}>
							{row.getVisibleCells().map((cell, index) => (
								<td key={cell.id}>
									{flexRender(
										cell.column.columnDef.cell,
										cell.getContext(),
									)}
								</td>
							))}
						</tr>
					))}
				</tbody>
			</table>
		</div>
	);
};

export default NewTable;
