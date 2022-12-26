import { HeartIcon, Bars3Icon } from "@heroicons/react/24/solid";
import { Menu, Transition } from "@headlessui/react";
import { Fragment } from "react";
import PropTypes from "prop-types";

/**
 * Represent a skin to be displayed in a grid.
 */
function SkinCard({ description }) {
	return (
		<div className="rounded h-72 w-56 border shadow shadow-slate-500 bg-blue-100">
			<div className="h-3/4">Skin Here</div>
			<div className="h-1/4 px-4 py-2 border-t border-t-slate-500">
				<h2 className="font-bold text-clip text-lg">{description}</h2>
				<h3>by AFCM</h3>
				<div className="top-2 static right-2 flex align-baseline items-end">
					<button className="h-8 w-8">
						<HeartIcon className="text-red-500" />
					</button>
					<Menu as="div" className="relative ml-3">
						<div>
							<Menu.Button className="flex items-center text-sm h-8 w-8">
								<span className="sr-only">Skin Options</span>
								<Bars3Icon className="" />
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
									<a
										href={"google.com"}
										className={
											"block px-4 py-2 text-sm text-gray-700"
										}
									>
										{"Download Skin"}
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
