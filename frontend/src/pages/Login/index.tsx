import { useState, useEffect } from "react";
import { Navigate, useNavigate } from "react-router-dom";
import { useRecoilValue } from "recoil";
import axios from "axios";
import { ExclamationTriangleIcon } from "@heroicons/react/24/solid";
import { AuthStateIsAuthenticated } from "../../states/auth";
import ApiUrls from "../../services/api_urls";

function Login() {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	const [canSubmit, setCanSubmit] = useState(false);

	const [loading, setLoading] = useState(false);
	const [err, setErr] = useState(null);

	useEffect(() => {
		if (email === "" || password === "") {
			setCanSubmit(false);
		} else {
			setCanSubmit(true);
		}
	}, [email, password, setCanSubmit]);

	const navigate = useNavigate();

	// Navigate to homepage if already authenticated
	if (useRecoilValue(AuthStateIsAuthenticated)) {
		return <Navigate to="/" />;
	}

	/**
	 * @param {React.FormEvent<HTMLFormElement>} e
	 */
	async function handleSubmit(e) {
		e.preventDefault();
		setLoading(true);
		await axios
			.post(ApiUrls.AccountLogin, {
				email: email,
				password: password,
			})
			.then((r) => {
				//console.log(r);
				setLoading(false);
				navigate("/");
			})
			.catch((e) => {
				//console.log(e);
				setLoading(false);
				setErr(e);
			});
	}

	return (
		<div className="flex justify-center align-middle">
			<div className="app-pannel m-10 flex max-w-md flex-col p-6">
				<h1 className="panel-text-heading">Login</h1>
				<form
					className="mt-4 flex flex-col gap-4"
					onSubmit={handleSubmit}
				>
					<label className="block">
						<span className="text-slate-800">Email</span>
						<input
							type="email"
							name="email"
							id="email"
							placeholder="someone@gmail.com"
							className="form"
							value={email}
							onChange={(e) => {
								setEmail(e.target.value);
							}}
						/>
					</label>
					<label className="block">
						<span className="text-slate-800">Password</span>
						<input
							type="password"
							name="password"
							id="password"
							placeholder="SuperStrongPassword"
							className="form"
							value={password}
							onChange={(e) => {
								setPassword(e.target.value);
							}}
						/>
					</label>
					{err !== null ? (
						<div className="mt-4 flex flex-row items-center gap-2 text-red-600">
							<ExclamationTriangleIcon className="h-8 w-8" />
							{err.message}
						</div>
					) : undefined}
					<div className="mt-4 flex gap-4">
						<button
							className="button-secondary w-full"
							onClick={(e) => {
								navigate("/register");
							}}
						>
							Register
						</button>
						<button
							className={`button-primary w-full disabled:cursor-not-allowed disabled:bg-black ${
								loading ? "cursor-wait" : null
							}`}
							type="submit"
							disabled={!canSubmit}
						>
							Login
						</button>
					</div>
				</form>
			</div>
		</div>
	);
}

export default Login;
