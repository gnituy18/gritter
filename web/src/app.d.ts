declare namespace App {
  type User = import("$types").User;

  interface Locals {
    currentUser?: User;
    sessionId?: string;
  }

  interface Platform {}

  interface Session {
    currentUser?: User;
    sessionId?: string;
  }

  interface Stuff {}
}

interface ImportMetaEnv {
  ENV_GOOGLE_CLIENT_ID: string;
  ENV_BACKEND_HOST: string;
}
