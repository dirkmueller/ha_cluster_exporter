package crmmon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p := NewCrmMonParser("foo")
	assert.Equal(t, "foo", p.crmMonPath)
}

func TestParse(t *testing.T) {
	p := NewCrmMonParser("../../../test/fake_crm_mon.sh")
	data, err := p.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "2.0.0", data.Version)
	assert.Equal(t, 8, data.Summary.Resources.Number)
	assert.Equal(t, 1, data.Summary.Resources.Disabled)
	assert.Equal(t, 0, data.Summary.Resources.Blocked)
	assert.Equal(t, "Fri Oct 18 11:48:22 2019", data.Summary.LastChange.Time)
	assert.Equal(t, 2, data.Summary.Nodes.Number)
	assert.Equal(t, "node01", data.Nodes[0].Name)
	assert.Equal(t, "1084783375", data.Nodes[0].Id)
	assert.Equal(t, true, data.Nodes[0].Online)
	assert.Equal(t, true, data.Nodes[0].ExpectedUp)
	assert.Equal(t, true, data.Nodes[0].DC)
	assert.Equal(t, false, data.Nodes[0].Unclean)
	assert.Equal(t, false, data.Nodes[0].Shutdown)
	assert.Equal(t, false, data.Nodes[0].StandbyOnFail)
	assert.Equal(t, false, data.Nodes[0].Maintenance)
	assert.Equal(t, false, data.Nodes[0].Pending)
	assert.Equal(t, false, data.Nodes[0].Standby)
	assert.Equal(t, "node02", data.Nodes[1].Name)
	assert.Equal(t, "1084783376", data.Nodes[1].Id)
	assert.Equal(t, true, data.Nodes[1].Online)
	assert.Equal(t, true, data.Nodes[1].ExpectedUp)
	assert.Equal(t, false, data.Nodes[1].DC)
	assert.Equal(t, false, data.Nodes[1].Unclean)
	assert.Equal(t, false, data.Nodes[1].Shutdown)
	assert.Equal(t, false, data.Nodes[1].StandbyOnFail)
	assert.Equal(t, false, data.Nodes[1].Maintenance)
	assert.Equal(t, false, data.Nodes[1].Pending)
	assert.Equal(t, false, data.Nodes[1].Standby)
	assert.Equal(t, "node01", data.NodeHistory.Nodes[0].Name)
	assert.Equal(t, 5000, data.NodeHistory.Nodes[0].ResourceHistory[0].MigrationThreshold)
	assert.Equal(t, 2, data.NodeHistory.Nodes[0].ResourceHistory[1].FailCount)
	assert.Equal(t, "rsc_SAPHana_PRD_HDB00", data.NodeHistory.Nodes[0].ResourceHistory[0].Name)
	assert.Equal(t, 4, len(data.Resources))
	assert.Equal(t, "test-stop", data.Resources[0].Id)
	assert.Equal(t, false, data.Resources[0].Active)
	assert.Equal(t, "Stopped", data.Resources[0].Role)
}

func TestParseClones(t *testing.T) {
	p := NewCrmMonParser("../../../test/fake_crm_mon.sh")
	data, err := p.Parse()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(data.Clones))
	assert.Equal(t, "msl_SAPHana_PRD_HDB00", data.Clones[0].Id)
	assert.Equal(t, "cln_SAPHanaTopology_PRD_HDB00", data.Clones[1].Id)
	assert.Equal(t, 2, len(data.Clones[0].Resources))
	assert.Equal(t, 2, len(data.Clones[1].Resources))
	assert.Equal(t, "rsc_SAPHana_PRD_HDB00", data.Clones[0].Resources[0].Id)
	assert.Equal(t, "Master", data.Clones[0].Resources[0].Role)
	assert.Equal(t, "rsc_SAPHana_PRD_HDB00", data.Clones[0].Resources[1].Id)
	assert.Equal(t, "Slave", data.Clones[0].Resources[1].Role)
}

func TestParseGroups(t *testing.T) {
	p := NewCrmMonParser("../../../test/fake_crm_mon.sh")
	data, err := p.Parse()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(data.Groups))

	assert.Equal(t, "grp_HA1_ASCS00", data.Groups[0].Id)
	assert.Equal(t, 3, len(data.Groups[0].Resources))
	assert.Equal(t, "rsc_ip_HA1_ASCS00", data.Groups[0].Resources[0].Id)
	assert.Equal(t, "rsc_fs_HA1_ASCS00", data.Groups[0].Resources[1].Id)
	assert.Equal(t, "rsc_sap_HA1_ASCS00", data.Groups[0].Resources[2].Id)

	assert.Equal(t, "grp_HA1_ERS10", data.Groups[1].Id)
	assert.Equal(t, 3, len(data.Groups[1].Resources))
	assert.Equal(t, "rsc_ip_HA1_ERS10", data.Groups[1].Resources[0].Id)
	assert.Equal(t, "rsc_fs_HA1_ERS10", data.Groups[1].Resources[1].Id)
	assert.Equal(t, "rsc_sap_HA1_ERS10", data.Groups[1].Resources[2].Id)
}

func TestParseNodeAttributes(t *testing.T) {
	p := NewCrmMonParser("../../../test/fake_crm_mon.sh")
	data, err := p.Parse()
	assert.NoError(t, err)
	assert.Len(t, data.NodeAttributes.Nodes, 2)
	assert.Equal(t, "node01", data.NodeAttributes.Nodes[0].Name)
	assert.Equal(t, "node02", data.NodeAttributes.Nodes[1].Name)

	assert.Len(t, data.NodeAttributes.Nodes[0].Attributes, 11)
	assert.Equal(t, "hana_prd_clone_state", data.NodeAttributes.Nodes[0].Attributes[0].Name)
	assert.Equal(t, "hana_prd_op_mode", data.NodeAttributes.Nodes[0].Attributes[1].Name)
	assert.Equal(t, "hana_prd_remoteHost", data.NodeAttributes.Nodes[0].Attributes[2].Name)
	assert.Equal(t, "hana_prd_roles", data.NodeAttributes.Nodes[0].Attributes[3].Name)
	assert.Equal(t, "hana_prd_site", data.NodeAttributes.Nodes[0].Attributes[4].Name)
	assert.Equal(t, "hana_prd_srmode", data.NodeAttributes.Nodes[0].Attributes[5].Name)
	assert.Equal(t, "hana_prd_sync_state", data.NodeAttributes.Nodes[0].Attributes[6].Name)
	assert.Equal(t, "hana_prd_version", data.NodeAttributes.Nodes[0].Attributes[7].Name)
	assert.Equal(t, "hana_prd_vhost", data.NodeAttributes.Nodes[0].Attributes[8].Name)
	assert.Equal(t, "lpa_prd_lpt", data.NodeAttributes.Nodes[0].Attributes[9].Name)
	assert.Equal(t, "master-rsc_SAPHana_PRD_HDB00", data.NodeAttributes.Nodes[0].Attributes[10].Name)

	assert.Equal(t, "PROMOTED", data.NodeAttributes.Nodes[0].Attributes[0].Value)
	assert.Equal(t, "logreplay", data.NodeAttributes.Nodes[0].Attributes[1].Value)
	assert.Equal(t, "node02", data.NodeAttributes.Nodes[0].Attributes[2].Value)
	assert.Equal(t, "4:P:master1:master:worker:master", data.NodeAttributes.Nodes[0].Attributes[3].Value)
	assert.Equal(t, "PRIMARY_SITE_NAME", data.NodeAttributes.Nodes[0].Attributes[4].Value)
	assert.Equal(t, "sync", data.NodeAttributes.Nodes[0].Attributes[5].Value)
	assert.Equal(t, "PRIM", data.NodeAttributes.Nodes[0].Attributes[6].Value)
	assert.Equal(t, "2.00.040.00.1553674765", data.NodeAttributes.Nodes[0].Attributes[7].Value)
	assert.Equal(t, "node01", data.NodeAttributes.Nodes[0].Attributes[8].Value)
	assert.Equal(t, "1571392102", data.NodeAttributes.Nodes[0].Attributes[9].Value)
	assert.Equal(t, "150", data.NodeAttributes.Nodes[0].Attributes[10].Value)

	assert.Len(t, data.NodeAttributes.Nodes[1].Attributes, 11)
	assert.Equal(t, "hana_prd_clone_state", data.NodeAttributes.Nodes[0].Attributes[0].Name)
	assert.Equal(t, "hana_prd_op_mode", data.NodeAttributes.Nodes[0].Attributes[1].Name)
	assert.Equal(t, "hana_prd_remoteHost", data.NodeAttributes.Nodes[0].Attributes[2].Name)
	assert.Equal(t, "hana_prd_roles", data.NodeAttributes.Nodes[0].Attributes[3].Name)
	assert.Equal(t, "hana_prd_site", data.NodeAttributes.Nodes[0].Attributes[4].Name)
	assert.Equal(t, "hana_prd_srmode", data.NodeAttributes.Nodes[0].Attributes[5].Name)
	assert.Equal(t, "hana_prd_sync_state", data.NodeAttributes.Nodes[0].Attributes[6].Name)
	assert.Equal(t, "hana_prd_version", data.NodeAttributes.Nodes[0].Attributes[7].Name)
	assert.Equal(t, "hana_prd_vhost", data.NodeAttributes.Nodes[0].Attributes[8].Name)
	assert.Equal(t, "lpa_prd_lpt", data.NodeAttributes.Nodes[0].Attributes[9].Name)
	assert.Equal(t, "master-rsc_SAPHana_PRD_HDB00", data.NodeAttributes.Nodes[0].Attributes[10].Name)

	assert.Equal(t, "DEMOTED", data.NodeAttributes.Nodes[1].Attributes[0].Value)
	assert.Equal(t, "logreplay", data.NodeAttributes.Nodes[1].Attributes[1].Value)
	assert.Equal(t, "node01", data.NodeAttributes.Nodes[1].Attributes[2].Value)
	assert.Equal(t, "4:S:master1:master:worker:master", data.NodeAttributes.Nodes[1].Attributes[3].Value)
	assert.Equal(t, "SECONDARY_SITE_NAME", data.NodeAttributes.Nodes[1].Attributes[4].Value)
	assert.Equal(t, "sync", data.NodeAttributes.Nodes[1].Attributes[5].Value)
	assert.Equal(t, "SOK", data.NodeAttributes.Nodes[1].Attributes[6].Value)
	assert.Equal(t, "2.00.040.00.1553674765", data.NodeAttributes.Nodes[1].Attributes[7].Value)
	assert.Equal(t, "node02", data.NodeAttributes.Nodes[1].Attributes[8].Value)
	assert.Equal(t, "30", data.NodeAttributes.Nodes[1].Attributes[9].Value)
	assert.Equal(t, "100", data.NodeAttributes.Nodes[1].Attributes[10].Value)
}
