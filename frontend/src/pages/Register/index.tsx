import React, { useContext, useEffect, useState } from "react";
import { Navigate, useNavigate, Link } from "react-router-dom";
import axios from "axios";
import { ExclamationTriangleIcon } from "@heroicons/react/24/solid";
import ApiUrls from "../../services/api_urls.ts";
import { AppContext } from "../../services/AppContext.tsx";

function Register() {
    const { loggedIn } = useContext(AppContext);

    // Navigate to homepage if already authenticated
    if (loggedIn) {
        return <Navigate to="/" />;
    }

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [agreedTOS, setAgreedTOS] = useState(false);

    const [canSubmit, setCanSubmit] = useState(false);

    const [loading, setLoading] = useState(false);
    const [err, setErr] = useState<string | null>(null);

    useEffect(() => {
        if (username === "" || password === "" || !agreedTOS) {
            setCanSubmit(false);
        } else {
            setCanSubmit(true);
        }
    }, [username, password, agreedTOS, setCanSubmit]);

    const navigate = useNavigate();

    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        setLoading(true);
        await axios
            .post(ApiUrls.AccountRegister, {
                username: username,
                password: password,
            })
            .then(() => {
                //console.log(r);
                setLoading(false);
                navigate("/login");
            })
            .catch((e: { message: string }) => {
                console.log(e);
                setLoading(false);
                setErr(e.message);
            });
    }

    return (
        <div className="flex justify-center align-middle">
            <div className="app-pannel mx-0 my-10 flex w-full max-w-md flex-col p-6 md:mx-10">
                <h1 className="panel-text-heading">Register</h1>
                <form className="mt-4 flex flex-col" onSubmit={handleSubmit}>
                    <label htmlFor="username" className="mb-2 block select-none text-sm font-medium dark:text-white">
                        Username
                    </label>
                    <input
                        type="text"
                        name="username"
                        id="username"
                        spellCheck={false}
                        className="form-input2"
                        onInput={(e) => setUsername((e.target as HTMLInputElement).value)}
                    />

                    <label
                        htmlFor="password"
                        className="mb-2 mt-4 block select-none text-sm font-medium dark:text-white"
                    >
                        Password
                    </label>
                    <input
                        type="password"
                        name="password"
                        id="password"
                        placeholder=""
                        className="form-input2"
                        onInput={(e) => setPassword((e.target as HTMLInputElement).value)}
                    />

                    <div className="mt-4 flex">
                        <input
                            type="checkbox"
                            name="understand"
                            id="understand"
                            className="form-checkbox"
                            defaultChecked={false}
                            onChange={(e) => setAgreedTOS(e.target.checked)}
                        />
                        <label
                            htmlFor="understand"
                            className="ms-3 select-none text-sm text-gray-500 dark:text-gray-400"
                        >
                            I aggree to the{" "}
                            <Link to="/" className={`text-blue-500 after:content-["_↗"]`}>
                                Terms of Service
                            </Link>
                        </label>
                    </div>
                    {err !== null ? (
                        <div className="mt-4 flex flex-row items-center gap-2 text-red-600">
                            <ExclamationTriangleIcon className="h-8 w-8" />
                            {err}
                        </div>
                    ) : undefined}
                    <button
                        className={`button-primary mt-4 w-full ${loading ? "cursor-wait" : null}`}
                        type={"submit"}
                        disabled={!canSubmit}
                    >
                        Register
                    </button>
                </form>
            </div>
        </div>
    );
}

export default Register;
