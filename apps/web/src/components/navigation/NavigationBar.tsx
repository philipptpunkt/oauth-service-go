import Link from "next/link";

interface NaviagtionItem {
  href: string;
  title: string;
}

const topSideNavigationItems = [
  {
    href: "/login",
    title: "Login",
  },
  {
    href: "/register",
    title: "Start free",
  },
];

function SideNavigationLink({ href, title }: NaviagtionItem) {
  return (
    <Link href={href} className="rounded-md border px-8 py-2 hover:bg-gray-400">
      {title}
    </Link>
  );
}

export function NavigationBar() {
  return (
    <div className="sticky top-0 flex h-16 items-center justify-between bg-blue-200 px-4">
      <Link href="/" className="bg-blue-700 px-8 py-4 text-white">
        Blue Auth
      </Link>
      <div>Main Naviagtion</div>
      <div>
        {topSideNavigationItems.map((link, index) => (
          <SideNavigationLink key={index} {...link} />
        ))}
      </div>
    </div>
  );
}
