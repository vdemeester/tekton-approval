import { Component } from '@angular/core';

export interface ApprovalElement {
  name: string;
  url: string;
  status: Status;
}

export enum Status {
  UNKNOWN = 'UNKNOWN',
  APPROVED = 'APPROVED',
  DECLINED = 'DECLINED'
}

const ELEMENT_DATA: ApprovalElement[] = [
  { url: 'http://test.com', name: 'Hydrogen', status: Status.UNKNOWN },
  { url: 'http://test.com', name: 'Helium', status: Status.UNKNOWN },
  { url: 'http://test.com', name: 'Lithium', status: Status.UNKNOWN },
  { url: 'http://test.com', name: 'Beryllium', status: Status.APPROVED },
  { url: 'http://test.com', name: 'Boron', status: Status.DECLINED },
  { url: 'http://test.com', name: 'Carbon', status: Status.APPROVED },
  { url: 'http://test.com', name: 'Nitrogen', status: Status.DECLINED },
  { url: 'http://test.com', name: 'Oxygen', status: Status.APPROVED },
  { url: 'http://test.com', name: 'Fluorine', status: Status.DECLINED },
  { url: 'http://test.com', name: 'Neon', status: Status.DECLINED },
];

@Component({
  selector: 'app-tekton-approvals',
  templateUrl: './tekton-approvals.component.html',
  styleUrls: ['./tekton-approvals.component.scss']
})
export class TektonApprovalsComponent {

  displayedColumns: string[] = ['name', 'url', 'status'];
  dataSource = ELEMENT_DATA;

}
