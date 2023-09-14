const emailRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;

const testPattern = (pattern: string): boolean => {
	return emailRegex.test(pattern);
};

export default testPattern;
