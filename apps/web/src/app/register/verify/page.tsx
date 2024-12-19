export default function EmailVerificationPage() {
  return (
    <div className="flex h-screen flex-col items-center justify-center">
      <h1 className="mb-16">Thanks mate</h1>
      <p>Please enter the code from the email below</p>
      <input placeholder="Code" />
      <p>
        If you have not received an Email with the code you can request a new
        code here
      </p>
      <button>Request new Code</button>
    </div>
  );
}
