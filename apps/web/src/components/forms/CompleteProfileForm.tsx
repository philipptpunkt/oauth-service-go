export function CompleteProfileForm() {
  return (
    <div>
      <form>
        <div>
          <input type="text" id="firstname" placeholder="First name" />
        </div>
        <div>
          <input type="text" id="lastname" placeholder="Last name" />
        </div>
        <div>
          <input type="text" id="jobTitle" placeholder="Job title" />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}
