function Button({className, textClassName, text}) {
  return (
    <>
	    <button className={className}>
	    	<span className={textClassName}>{text}</span>
	    </button>
    </>
  )
}

export default Button;
