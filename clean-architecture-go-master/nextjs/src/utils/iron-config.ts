const ironConfig = {
    password: process.env.COOKIE_KEY as string,
    cookieName: "fullcycle-session",
    cookieOptions: {
        // the next line allows to use the session in non-https environments like
        // Next.js dev mode (http://localhost:3000)
        secure: process.env.NODE_ENV === "production",
    },
};

declare module "iron-session" {
    interface IronSessionData {
        account?: {
            id: number;
            name: string;
            token: string;
        };
    }
}

export default ironConfig;