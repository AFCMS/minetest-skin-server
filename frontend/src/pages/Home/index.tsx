import SkinCard from "../../components/SkinCard";

function Home() {
    return (
            <div className="justify-center align-middle flex">
                <div className="max-w-screen-xl m-10 app-pannel flex flex-col gap-8 container">
                    <div className="w-full h-24 bg-slate-400 rounded-t px-8"></div>
                    <div className="grid gap-8 px-8 pb-8 grid-cols-1 md:grid-cols-4 max xl:grid-cols-5 grid-flow-row justify-items-center">
                        {(() => {
                            let e = [];

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
                </div>
            </div>
    );
}

export default Home;
