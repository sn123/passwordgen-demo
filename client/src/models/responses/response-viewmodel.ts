export class ResponseViewModel<T> {
  public data: T | null = null;

  public static hasErrors<T>(viewModel: ResponseViewModel<T>): boolean {
    return !viewModel || !viewModel.data;
  }

  public static getNullable<T>(
    request: ResponseViewModel<T>
  ): ResponseViewModel<T> {
    return request ? request : new ResponseViewModel();
  }
}
