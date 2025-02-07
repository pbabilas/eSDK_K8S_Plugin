/*
 *  Copyright (c) Huawei Technologies Co., Ltd. 2020-2022. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package driver

import (
	"strings"

	"huawei-csi-driver/utils/k8sutils"
)

type Driver struct {
	name              string
	version           string
	useMultiPath      bool
	scsiMultiPathType string
	nvmeMultiPathType string
	k8sUtils          k8sutils.Interface
	nodeName          string
}

func NewDriver(name, version string, useMultiPath bool, scsiMultiPathType, nvmeMultiPathType string,
	k8sUtils k8sutils.Interface, nodeName string) *Driver {
	return &Driver{
		name:              name,
		version:           version,
		useMultiPath:      useMultiPath,
		scsiMultiPathType: scsiMultiPathType,
		nvmeMultiPathType: nvmeMultiPathType,
		k8sUtils:          k8sUtils,
		nodeName:          strings.TrimSpace(nodeName),
	}
}
