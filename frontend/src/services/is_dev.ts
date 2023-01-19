const development = process.env.NODE_ENV === "development";

/**
 * @returns Is running in development mode
 */
const isDev = (): boolean => development;

export default isDev;
