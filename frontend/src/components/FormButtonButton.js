function FormButtonButton({className, textClassName, text, id}) {
  return (
    <>
	    <button className={className} id={id}>
	    	<span className={textClassName}>{text}</span>
	    </button>
    </>
  )
}

export default FormButtonButton;
