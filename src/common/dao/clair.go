// copyright (c) 2017 vmware, inc. all rights reserved.
//
// licensed under the apache license, version 2.0 (the "license");
// you may not use this file except in compliance with the license.
// you may obtain a copy of the license at
//
//    http://www.apache.org/licenses/license-2.0
//
// unless required by applicable law or agreed to in writing, software
// distributed under the license is distributed on an "as is" basis,
// without warranties or conditions of any kind, either express or implied.
// see the license for the specific language governing permissions and
// limitations under the license.

package dao

import (
	"github.com/vmware/harbor/src/common/models"

	"fmt"
	"time"
)

//SetClairVulnTimestamp update the last_update of a namespace. If there's no record for this namespace, one will be created.
func SetClairVulnTimestamp(namespace string, timestamp time.Time) error {
	o := GetOrmer()
	rec := &models.ClairVulnTimestamp{
		Namespace:  namespace,
		LastUpdate: timestamp,
	}
	created, _, err := o.ReadOrCreate(rec, "Namespace")
	if err != nil {
		return err
	}
	if !created {
		rec.LastUpdate = timestamp
		n, err := o.Update(rec)
		if err != nil {
			return err
		}
		if n == 0 {
			return fmt.Errorf("No record is updated, record: %v", *rec)
		}
	}
	return nil
}

//ListClairVulnTimestamps return a list of all records in vuln timestamp table.
func ListClairVulnTimestamps() ([]*models.ClairVulnTimestamp, error) {
	var res []*models.ClairVulnTimestamp
	o := GetOrmer()
	_, err := o.QueryTable(models.ClairVulnTimestampTable).All(&res)
	return res, err
}
