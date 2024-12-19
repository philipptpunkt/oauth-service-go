import { Card } from "@oauth-service-go/ui/Cards";
import { RegisterClientForm } from "../../components/forms/RegisterClientForm";

export default function RegisterClientPage() {
  return (
    <div className="flex h-screen items-center">
      <div className="w-1/2 p-16">
        <h1>Start for free. Easy to use. Create your account now</h1>
      </div>
      <Card className="w-1/2 max-w-[500px]">
        <h2>Register</h2>
        <RegisterClientForm />
      </Card>
    </div>
  );
}
