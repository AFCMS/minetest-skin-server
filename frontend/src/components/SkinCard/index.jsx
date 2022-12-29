import {
	HeartIcon,
	Bars3Icon,
	FolderArrowDownIcon,
	DocumentArrowDownIcon,
} from "@heroicons/react/24/solid";
import { Suspense } from "react";
import { Menu, Transition } from "@headlessui/react";
import { Canvas } from "@react-three/fiber";
import { PerspectiveCamera, OrbitControls } from "@react-three/drei";
//import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import { Fragment } from "react";
import PropTypes from "prop-types";
import SkinModel from "../SkinModel";

/**
 * Represent a skin to be displayed in a grid.
 * @param {{description: string}} param0
 */
function SkinCard({ description }) {
	//const gltf = useGLTF(skinModel, true);
	//const texture = useTexture(skinTexture);

	//const group = useRef();

	return (
		<div className="h-72 w-56 rounded border bg-blue-100 shadow shadow-slate-500">
			<div className="h-3/4">
				<Canvas className="bg-scroll">
					<Suspense fallback={null}>
						<PerspectiveCamera />
						<OrbitControls />
						<ambientLight intensity={0.5} />
						<directionalLight />
						<SkinModel />
						{/*<Environment preset="sunset" background />*/}
					</Suspense>
				</Canvas>
			</div>
			<div className="relative h-1/4 border-t border-t-slate-500 px-4 py-2">
				<h2 className="text-clip text-lg font-bold text-slate-800">
					{description}
				</h2>
				<h3 className="text-slate-800">by AFCM</h3>
				<div className="absolute bottom-2 right-2 flex items-end align-baseline">
					<button aria-label="like skin" className="h-8 w-8">
						<HeartIcon className="text-red-500" />
					</button>
					<Menu as="div" className="relative ml-3">
						<div>
							<Menu.Button
								className="flex h-8 w-8 items-center text-sm"
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
											"flex h-8 cursor-pointer flex-row gap-2 p-1 px-2 align-baseline text-sm text-slate-800 hover:bg-slate-200"
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
											"flex h-8 flex-row gap-2 p-1 px-2 align-baseline text-sm text-slate-800 hover:bg-slate-200"
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
											"flex h-8 flex-row gap-2 p-1 px-2 align-baseline text-sm text-slate-800 hover:bg-slate-200"
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
