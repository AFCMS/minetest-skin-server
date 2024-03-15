import { CodeBracketIcon } from "@heroicons/react/24/outline";

function About() {
    return (
            <div className="flex justify-center align-middle">
                <div className=" container m-10 max-w-prose rounded bg-slate-50 p-8 shadow-md">
                    <h1 className="panel-text-heading">About</h1>
                    <p className="mt-2">
                        Minetest Skin Server is a service to upload Minetest skins
                        and let servers retreive them. It is licensed under GPLv3.
                    </p>
                    <a
                            className="panel-text-link"
                            href="https://github.com/AFCMS/minetest-skin-server"
                    >
                        Source Code
                    </a>
                    <br/>
                    <section className="my-4 ml-6">
                        <ul className="flex flex-col">
                            <li className="align-baseline text-lg font-medium text-slate-800">
                                <CodeBracketIcon className="mr-2 inline h-8"/>
                                Licensed under GPLv3
                            </li>
                            <li className="align-baseline text-lg font-medium text-slate-800">
                                <CodeBracketIcon className="mr-2 inline h-8"/>
                                Supports MineClone2, skindb, etc
                            </li>
                        </ul>
                    </section>
                </div>
            </div>
    );
}

export default About;
