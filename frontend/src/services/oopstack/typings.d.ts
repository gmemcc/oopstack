declare namespace API {
  type ErrorResponse = {
    /** Error message */
    message?: string;
  };

  type GreetingResponse = {
    /** The personalized greeting message.
Example: Hello Alex */
    greeting?: string;
  };

  type NameRequest = {
    /** Name of the person to greet.
Required: true */
    name?: string;
  };
}
