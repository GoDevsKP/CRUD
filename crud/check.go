package crud


func Check(query []string) bool {

	if (len(query) == 0) {
		return false;
	}
	action := query[0]

	switch action{
		case "create":
			return CreateValidator(query);
		case "select":
			return SelectValidator(query)
		case "delete":
			return DeleteValidator(query)
		case "update":
			return UpdateValidator(query)
		default:
			return false

	}
}

