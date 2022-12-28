const development = process.env.NODE_ENV === "development";

/**
 * @returns {boolean} Is running in development mode
 */
const isDev = () => {
	return development;
};

export default isDev;
