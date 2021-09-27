package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	DefaultPageLimit = 50

	flagPageOffset = "offset"
	flagPageLimit  = "limit"
)

// PageParams keeps query pagination params.
type PageParams struct {
	Offset int
	Limit  int
}

// NewDefaultPageParams creates a new PageParams with default params.
func NewDefaultPageParams() PageParams {
	return PageParams{
		Limit: DefaultPageLimit,
	}
}

// Validate perform a basic object validation.
func (p PageParams) Validate() error {
	if p.Offset < 0 {
		return fmt.Errorf("offset: must be GTE 0")
	}
	if p.Limit <= 0 {
		return fmt.Errorf("limit: must be GTE 1")
	}

	return nil
}

// AddPageParamsToCmd adds pagination flag to cobra.Command.
func AddPageParamsToCmd(cmd *cobra.Command) {
	cmd.Flags().Int(flagPageOffset, 0, "Query page offset")
	cmd.Flags().Int(flagPageLimit, 50, "Query page limit")
}

// ParsePageParams builds PageParams from cobra.Command flags.
func ParsePageParams(cmd *cobra.Command) (PageParams, error) {
	offset, err := cmd.Flags().GetInt(flagPageOffset)
	if err != nil {
		return PageParams{}, fmt.Errorf("reading %q parameter: %w", flagPageOffset, err)
	}

	limit, err := cmd.Flags().GetInt(flagPageLimit)
	if err != nil {
		return PageParams{}, fmt.Errorf("reading %q parameter: %w", flagPageLimit, err)
	}

	return PageParams{
		Offset: offset,
		Limit:  limit,
	}, nil
}
