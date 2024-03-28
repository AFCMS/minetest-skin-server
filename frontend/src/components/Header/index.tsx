import { Fragment, useContext } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import { Disclosure, Menu, Transition } from "@headlessui/react";
import { Bars3Icon, BellIcon, XMarkIcon } from "@heroicons/react/24/outline";
import minetestIcon from "../../assets/minetest_logo.png";
import profileImagePlaceholder from "../../assets/character_base_head.png";
import { AppContext } from "../../services/AppContext.tsx";

const navigation: { name: string; href: string }[] = [
    { name: "Home", href: "/" },
    { name: "Search", href: "/search" },
    { name: "About", href: "/about" },
];

function Header(): JSX.Element {
    const { loggedIn, logout, loadingUser, username } = useContext(AppContext);

    const navigate = useNavigate();

    return (
        <>
            <div className="min-h-full">
                <Disclosure as="nav" className="bg-slate-600">
                    {({ open }) => (
                        <>
                            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                                <div className="flex h-16 items-center justify-between">
                                    <div className="flex items-center">
                                        <div className="flex-shrink-0">
                                            <img className="h-8 w-8" src={minetestIcon} alt="Minetest Skin Server" />
                                        </div>
                                        <div className="hidden md:block">
                                            <div className="ml-10 flex items-baseline space-x-4">
                                                {navigation.map((item) => (
                                                    <Disclosure.Button
                                                        as={NavLink}
                                                        key={item.name}
                                                        to={item.href}
                                                        className={(props: { isActive: boolean }) => {
                                                            return props.isActive
                                                                ? "button-navbar-active"
                                                                : "button-navbar";
                                                        }}
                                                    >
                                                        {item.name}
                                                    </Disclosure.Button>
                                                ))}
                                            </div>
                                        </div>
                                    </div>
                                    <div className="hidden md:block">
                                        <div className="ml-4 flex items-center md:ml-6">
                                            <button
                                                type="button"
                                                className="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white"
                                                aria-label="View notifications"
                                            >
                                                <BellIcon className="h-6 w-6" />
                                            </button>

                                            {/* Profile dropdown */}
                                            {loggedIn ? (
                                                <Menu as="div" className="relative ml-3">
                                                    <div>
                                                        <Menu.Button
                                                            className="flex max-w-xs items-center rounded bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-200"
                                                            aria-label="Open user menu"
                                                        >
                                                            <img
                                                                className="rendering-pixelated h-10 w-10 rounded"
                                                                src={profileImagePlaceholder}
                                                                alt=""
                                                            />
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
                                                            <Menu.Item key={"logout"}>
                                                                {({ active }) => (
                                                                    <button
                                                                        className={`${
                                                                            active ? "bg-slate-200" : ""
                                                                        } block w-full px-4 py-2 text-sm text-gray-700`}
                                                                        onClick={() => {
                                                                            logout().then(() => {
                                                                                navigate("/");
                                                                            });
                                                                        }}
                                                                    >
                                                                        Logout
                                                                    </button>
                                                                )}
                                                            </Menu.Item>
                                                        </Menu.Items>
                                                    </Transition>
                                                </Menu>
                                            ) : !loadingUser ? (
                                                <div className="relative ml-3 flex space-x-4">
                                                    <NavLink to="/register" className="button-secondary">
                                                        Register
                                                    </NavLink>
                                                    <NavLink to="/login" className="button-primary">
                                                        Login
                                                    </NavLink>
                                                </div>
                                            ) : undefined}
                                        </div>
                                    </div>
                                    <div className="-mr-2 flex md:hidden">
                                        {/* Mobile menu button */}
                                        <Disclosure.Button className="inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white">
                                            <span className="sr-only">Open main menu</span>
                                            {open ? (
                                                <XMarkIcon className="block h-6 w-6" aria-hidden="true" />
                                            ) : (
                                                <Bars3Icon className="block h-6 w-6" aria-hidden="true" />
                                            )}
                                        </Disclosure.Button>
                                    </div>
                                </div>
                            </div>

                            <Disclosure.Panel className="md:hidden">
                                <div className="space-y-1 px-2 pb-3 pt-2 sm:px-3">
                                    {navigation.map((item) => (
                                        <Disclosure.Button
                                            key={item.name}
                                            as={NavLink}
                                            to={item.href}
                                            className={(props: { isActive: boolean }) => {
                                                return props.isActive
                                                    ? "block rounded-md bg-gray-900 px-3 py-2 text-base font-medium text-white"
                                                    : "block rounded-md px-3 py-2 text-base font-medium text-gray-300 hover:bg-gray-700 hover:text-white";
                                            }}
                                        >
                                            {item.name}
                                        </Disclosure.Button>
                                    ))}
                                </div>
                                <div className="border-t border-gray-700 pb-3 pt-4">
                                    {loggedIn ? (
                                        <>
                                            <div className="flex items-center px-5">
                                                <div className="flex-shrink-0">
                                                    <img
                                                        className="rendering-pixelated h-10 w-10 rounded"
                                                        src={profileImagePlaceholder}
                                                        alt=""
                                                    />
                                                </div>
                                                <div className="ml-3">
                                                    <div className="text-base font-medium leading-none text-white">
                                                        {username}
                                                    </div>
                                                </div>
                                                <button
                                                    type="button"
                                                    className="ml-auto flex-shrink-0 rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                                                >
                                                    <span className="sr-only">View notifications</span>
                                                    <BellIcon className="h-6 w-6" aria-hidden="true" />
                                                </button>
                                            </div>
                                            <div className="mt-3 space-y-1 px-2">
                                                <Disclosure.Button
                                                    key={"logout"}
                                                    as={"button"}
                                                    className="block w-full rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
                                                    onClick={() => {
                                                        logout().then(() => {
                                                            navigate("/");
                                                        });
                                                    }}
                                                >
                                                    Logout
                                                </Disclosure.Button>
                                            </div>
                                        </>
                                    ) : (
                                        <div className="relative ml-3 flex gap-2">
                                            <Disclosure.Button as={NavLink} to="/register" className="button-secondary">
                                                Register
                                            </Disclosure.Button>
                                            <Disclosure.Button as={NavLink} to="/login" className="button-primary">
                                                Login
                                            </Disclosure.Button>
                                        </div>
                                    )}
                                </div>
                            </Disclosure.Panel>
                        </>
                    )}
                </Disclosure>
            </div>
        </>
    );
}

export default Header;
