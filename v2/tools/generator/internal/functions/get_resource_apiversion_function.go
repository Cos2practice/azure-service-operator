/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var GetAPIVersionFunctionName = "Get" + astmodel.APIVersionProperty

// NewGetAPIVersionFunction returns a function that returns a static API Version enum value ('APIVersion_Value')
func NewGetAPIVersionFunction(
	apiVersionTypeName astmodel.TypeName,
	apiVersionEnumValue astmodel.EnumValue,
	idFactory astmodel.IdentifierFactory,
) astmodel.Function {
	comment := fmt.Sprintf("returns the ARM API version of the resource. This is always %s", apiVersionEnumValue.Value)

	value := astbuilder.TextLiteral(apiVersionEnumValue.Value)
	result := NewObjectFunction(GetAPIVersionFunctionName, idFactory,
		createBodyReturningValue(
			value,
			astmodel.StringType,
			comment,
			ReceiverTypeStruct))

	result.AddPackageReference(astmodel.GenRuntimeReference)

	return result
}
