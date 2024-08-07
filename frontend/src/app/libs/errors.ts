export const getErrorMessage = (error: any): string => {
  if (error instanceof Error) {
    return error.message;
  } else if (typeof error === "string") {
    return error;
  } else if (typeof error === "object" && error.message) {
    return error.message;
  } else {
    return "An unknown error occurred.";
  }
};
