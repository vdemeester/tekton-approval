import { Component } from '@angular/core';
import { ApprovalsService } from 'src/app/approvals.service';

export interface ApprovalElement {
  name: string;
  url: string;
  status: Status;
}

export enum Status {
  UNKNOWN = 'Unknown',
  APPROVED = 'Approved',
  DECLINED = 'Declined'
}

@Component({
  selector: 'app-tekton-approvals',
  templateUrl: './tekton-approvals.component.html',
  styleUrls: ['./tekton-approvals.component.scss']
})
export class TektonApprovalsComponent {

  public dataSource: ApprovalElement[];
  public displayedColumns = ['name', 'url', 'status'];

  constructor(private approvalsService: ApprovalsService) {
    this.approvalsService.getApprovals().subscribe(data => this.dataSource = data);
  }

}
