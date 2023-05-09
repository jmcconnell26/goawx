package awx

type ExecutionEnvironmentsService struct {
	client *Client
}

type ExecutionEnvironmentsResponse struct {
	Pagination
	Results []*ExecutionEnvironment `json:"results"`
}

const executionEnvironmentsAPIEndpoint = "/api/v2/execution_environments/"

func (executionEnvironmentsService *ExecutionEnvironmentsService) ListExecutionEnvironments(params map[string]string) ([]*ExecutionEnvironment,
	*ExecutionEnvironmentsResponse,
	error) {
	result := new(ExecutionEnvironmentsResponse)
	response, err := executionEnvironmentsService.client.Requester.GetJSON(
		executionEnvironmentsAPIEndpoint,
		result,
		params)

	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(response)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
