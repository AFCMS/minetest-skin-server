/**
 * @returns Is running in development mode
 */
const isDev = (): boolean => import.meta.env.DEV;

export default isDev;
