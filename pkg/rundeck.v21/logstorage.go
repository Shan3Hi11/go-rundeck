package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// LogStorage represents log storage
type LogStorage responses.LogStorageResponse

// IncompleteLogStorage represents IncompleteLogStorage
type IncompleteLogStorage responses.IncompleteLogStorageResponse

// GetLogStorage gets the logstorage
func (c *Client) GetLogStorage() (*LogStorage, error) {
	if _, err := c.hasRequiredAPIVersion(17, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	ls := &LogStorage{}
	data, err := c.httpGet("system/logstorage", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ls); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return ls, nil
}

// GetIncompleteLogStorage gets executions with incomplete logstorage
func (c *Client) GetIncompleteLogStorage() (*IncompleteLogStorage, error) {
	if _, err := c.hasRequiredAPIVersion(14, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	ls := &IncompleteLogStorage{}
	data, err := c.httpGet("system/logstorage/incomplete", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ls); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return ls, nil
}

// ResumeIncompleteLogStorage resumes processing incomplete log storage uploads
// http://rundeck.org/docs/api/index.html#resume-incomplete-log-storage
func (c *Client) ResumeIncompleteLogStorage() (bool, error) {
	if _, err := c.hasRequiredAPIVersion(14, maxRundeckVersionInt); err != nil {
		return false, err
	}
	res := make(map[string]bool)
	data, err := c.httpPost("system/logstorage/incomplete/resume", requestJSON(), requestExpects(200))
	if err != nil {
		return false, err
	}
	if jsonErr := json.Unmarshal(data, &res); jsonErr != nil {
		return false, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	passed := res["resumed"]
	return passed, nil
}
