const clickOutside = (node: HTMLElement) => {
  const handleClick: (evt: Event) => void = (evt) => {
    if (node && !node.contains(evt.target as Node) && !evt.defaultPrevented) {
      node.dispatchEvent(new CustomEvent("clickOutside"));
    }
  };

  document.addEventListener("click", handleClick, true);

  return {
    destroy() {
      document.removeEventListener("click", handleClick, true);
    },
  };
};

export default clickOutside;
