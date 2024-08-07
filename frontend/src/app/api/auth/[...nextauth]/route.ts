import NextAuth, { AuthOptions, Session } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { getErrorMessage } from "@/app/libs/errors";
import { User } from "@/app/libs/types";

export const authOptions: AuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        try {
          const response = await fetch(`${process.env.API_URL}/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(credentials),
          });
          if (response.ok) {
            const user = await response.json();
            console.log(user);
            return user;
          }

          const badRequest = await response.json();
          console.log(badRequest);
          throw new Error(badRequest.error);
        } catch (error) {
          const message = getErrorMessage(error);
          throw new Error(message);
        }
      },
    }),
  ],
  session: {
    strategy: "jwt",
    maxAge: 60 * 60 * 24,
  },

  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.user = user;
      }
      return token;
    },
    async session({ session, token }) {
      let _session: Session | null = null;
      const user = token.user as User;
      if (user) {
        _session = { ...session, user: { ...user } };
      }
      return _session as Session;
    },
  },
  pages: {
    signIn: "/login",
    signOut: "/logout",
  },
  secret: process.env.NEXTAUTH_SECRET,
};

const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
