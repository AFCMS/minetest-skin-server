import {ArrowLeftIcon, ArrowRightIcon} from "@heroicons/react/24/solid";
import SkinCard from "../../components/SkinCard";

function SearchSkin() {
    return (
            <div className="justify-center align-middle flex">
                <div className="max-w-screen-xl m-10 app-pannel flex flex-col gap-8 container">
                    <div className="w-full h-24 bg-slate-400 rounded-t px-8"></div>
                    <div className="grid gap-8 px-8 pb-8 grid-cols-1 md:grid-cols-4 max xl:grid-cols-5 grid-flow-row justify-items-center">
                        {(() => {
                            const e = [];

                            for (let i = 0; i <= 10; i++) {
                                e.push(
                                        <SkinCard
                                                description={`A skin (${i})`}
                                                key={i}
                                        />
                                );
                            }

                            return e;
                        })()}
                    </div>
                    <nav className="place-content-center flex gap-4 m-6">
                        <button
                                aria-label="previous page"
                                className="rounded-md p-1 h-8 w-16 text-lg font-medium shadow-sm bg-blue-400 text-slate-50 flex justify-center"
                        >
                            <ArrowLeftIcon className="h-full"/>
                        </button>
                        <button
                                aria-label="next page"
                                className="rounded-md p-1 h-8 w-16 text-lg font-medium shadow-sm bg-blue-400 text-slate-50 flex justify-center"
                        >
                            <ArrowRightIcon className="h-full"/>
                        </button>
                    </nav>
                </div>
            </div>
    );
}

export default SearchSkin;
