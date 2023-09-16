type User = {
	id: string;
	email: string;
}

interface InputProps {
	type?: string;
}

declare namespace App {
	interface Locals {
		user: User | null;
	}


}

declare namespace svelteHTML {
	interface HTMLAttributes<T> {
		'on:clickOutside'?: CompositionEventHandler<T>;
	}
}