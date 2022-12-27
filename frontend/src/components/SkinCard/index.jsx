import {
	HeartIcon,
	Bars3Icon,
	FolderArrowDownIcon,
	DocumentArrowDownIcon,
} from "@heroicons/react/24/solid";
import { Menu, Transition } from "@headlessui/react";
import { Fragment } from "react";
import PropTypes from "prop-types";

/**
 * Represent a skin to be displayed in a grid.
 * @param {{description: string}} param0
 */
function SkinCard({ description }) {
	return (
		<div className="rounded h-72 w-56 border shadow shadow-slate-500 bg-blue-100">
			<div className="h-3/4">Skin Here</div>
			<div className="h-1/4 px-4 py-2 border-t border-t-slate-500 relative">
				<h2 className="font-bold text-clip text-lg text-slate-800">
					{description}
				</h2>
				<h3 className="text-slate-800">by AFCM</h3>
				<div className="bottom-2 absolute right-2 flex align-baseline items-end">
					<button aria-label="like skin" className="h-8 w-8">
						<HeartIcon className="text-red-500" />
					</button>
					<Menu as="div" className="relative ml-3">
						<div>
							<Menu.Button
								className="flex items-center text-sm h-8 w-8"
								aria-label="show options"
								aria-haspopup="listbox"
							>
								<span className="sr-only">Skin Options</span>
								<Bars3Icon className="text-slate-800" />
							</Menu.Button>
						</div>
						<Transition
							as={Fragment}
							enter="transition ease-out duration-100"
							enterFrom="transform opacity-0 scale-95"
							enterTo="transform opacity-100 scale-100"
							leave="transition ease-in duration-75"
							leaveFrom="transform opacity-100 scale-100"
							leaveTo="transform opacity-0 scale-95"
						>
							<Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
								<Menu.Item>
									<span
										className={
											"flex flex-row align-baseline h-8 p-1 gap-2 text-slate-800 text-sm px-2 hover:bg-slate-200"
										}
									>
										<DocumentArrowDownIcon className="" />
										Copy UUID
									</span>
								</Menu.Item>
								<Menu.Item>
									<a
										href={"https://google.com"}
										className={
											"flex flex-row align-baseline h-8 p-1 gap-2 text-slate-800 text-sm px-2 hover:bg-slate-200"
										}
									>
										<FolderArrowDownIcon className="" />
										Download Skin
									</a>
								</Menu.Item>
								<Menu.Item>
									<a
										href={"https://drive.google.com"}
										className={
											"flex flex-row align-baseline h-8 p-1 gap-2 text-slate-800 text-sm px-2 hover:bg-slate-200"
										}
									>
										<FolderArrowDownIcon className="" />
										Download Skin Head
									</a>
								</Menu.Item>
							</Menu.Items>
						</Transition>
					</Menu>
				</div>
			</div>
		</div>
	);
}

SkinCard.propTypes = {
	description: PropTypes.string.isRequired,
};

export default SkinCard;
