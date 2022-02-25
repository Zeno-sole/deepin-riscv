/*
 * Copyright (C) 2019 ~ 2021 Uniontech Software Technology Co.,Ltd
 *
 * Author:     dengbo <dengbo@uniontech.com>
 *
 * Maintainer: dengbo <dengbo@uniontech.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	info := DeviceInfo{
		Id:      111,
		Type:    1,
		Name:    "test",
		Enabled: true,
	}

	infos := DeviceInfos{
		&info,
	}

	assert.Equal(t, infos.Get(111).Id, info.Id)
}
