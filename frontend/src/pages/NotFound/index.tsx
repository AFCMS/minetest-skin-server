import { Link } from "react-router-dom";

// The 404, Not Found page
function NotFound() {
	return (
		<div className="flex justify-center align-middle">
			<div className=" container m-10 max-w-prose rounded bg-slate-50 p-8 shadow-md">
				<h1 className="panel-text-heading">Page Not Found</h1>
				<p className="mt-2">
					That page could not be found. The link may be broken, the
					page may have been deleted, or you may not have access to
					it.
				</p>
				<Link to="/" className="button-primary mt-4 inline-block">
					Back to Homepage
				</Link>
			</div>
		</div>
	);
}

export default NotFound;
