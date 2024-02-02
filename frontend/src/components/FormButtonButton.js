function FormButtonButton({className, textClassName, text, id, onClick}) {
  return (
    <>
	    <button className={className} id={id} type={id=="submit"?"submit":"button"} onClick={onClick}>
	    	<span className={textClassName}>{text}</span>
	    </button>
    </>
  )
}

export default FormButtonButton;
