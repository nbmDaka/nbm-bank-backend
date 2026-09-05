package validation

import (
	"testing"
)


func TestValidateUserID(t *testing.T) {


	tests := []struct{
		name string
		id int64
		wantError bool
	}{
		{
			name:"valid id",
			id:1,
			wantError:false,
		},
		{
			name:"zero id",
			id:0,
			wantError:true,
		},
		{
			name:"negative id",
			id:-1,
			wantError:true,
		},
	}



	for _, tt := range tests {

		t.Run(
			tt.name,
			func(t *testing.T){

				err := ValidateUserID(
					tt.id,
				)


				if tt.wantError && err == nil {

					t.Errorf(
						"expected error but got nil",
					)

				}


				if !tt.wantError && err != nil {

					t.Errorf(
						"unexpected error: %v",
						err,
					)

				}

			},
		)
	}
}